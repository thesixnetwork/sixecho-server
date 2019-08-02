import json
import os
import uuid
from datetime import datetime

import boto3
import mysql.connector
import numpy as np
import pycountry
from datasketch import LeanMinHash, MinHash, MinHashLSH
from sqlalchemy import create_engine
from sqlalchemy.orm import sessionmaker

from sixecho_model.category import Category
from sixecho_model.publisher import Publisher
import hmac
import hashlib

lambda_client = boto3.client('lambda')

ssm = boto3.client('ssm')

parameter = ssm.get_parameter(Name="SIXECHO_HOST_DB")
parameter2 = ssm.get_parameter(Name="SIXECHO_USER_DB")
parameter3 = ssm.get_parameter(Name="SIXECHO_PASSWORD_DB")

DB_HOST = parameter["Parameter"]["Value"]
DB_USER = parameter2["Parameter"]["Value"]
DB_PASSWORD = parameter3["Parameter"]["Value"]

URL_ENGINE = "mysql+mysqlconnector://%s:%s@%s/sixecho" % (DB_USER, DB_PASSWORD, DB_HOST)
ENGINE = create_engine(URL_ENGINE, echo=True)
Session = sessionmaker(bind=ENGINE)


def convert_str_to_minhash(digest):
    """Convert string that is including 128 numbers which to have a comma as middle between that numbers.
    Ex. 13241234,213242134,22342234,23423423,...,21341234 (128 numbers.)
    """
    data_array = np.array(digest.split(","), dtype=np.uint64)
    m1 = MinHash(hashvalues=data_array)
    return m1


def insert_mysql(api_key_id=None,
                 media_id=None,
                 digest=None,
                 sha256=None,
                 size_file=None,
                 meta_books=None):
    """Insert digital content to mysql for keeping information about the book.
    Args:
        api_key_id(string)  - Required  : api key from api gateway
        media_id(string)     - Required  : unique id we generate by time
        digest(string)      - Required  : digest from client
        sha256(string)      - Required  : sha256 generate from client
        size_file(string)   - Required  : size of file from client
    """
    mydb = mysql.connector.connect(host=DB_HOST,
                                   user=DB_USER,
                                   passwd=DB_PASSWORD,
                                   database="sixecho")
    mycursor = mydb.cursor()
    json_meta_books = json.dumps(meta_books)
    category_id = meta_books["category_id"]
    publisher_id = meta_books["publisher_id"]
    title = meta_books["title"]
    author = meta_books["author"]
    country_of_origin = meta_books["country_of_origin"]
    language = meta_books["language"]
    paperback = meta_books["paperback"]
    publish_date = meta_books["publish_date"]
    publish_date_str = datetime.utcfromtimestamp(publish_date).strftime(
        '%Y-%m-%d %H:%M:%S')

    sql = "INSERT INTO digital_contents (api_key_id, media_id, digest, sha256, size_file, meta_books,category_id,publisher_id,title,author,country_of_origin,language,paperback,publish_date) VALUES (%s, %s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s)"
    val = (api_key_id, media_id, digest, sha256, size_file, json_meta_books,
           category_id, publisher_id, title, author, country_of_origin,
           language, paperback, publish_date_str)
    mycursor.execute(sql, val)
    mydb.commit()
    mycursor.close()

def get_API_Secret(api_key):
    # import module
    import mysql.connector
    # import errorcode
    from mysql.connector import errorcode
    try:
        mydb = mysql.connector.connect(
            host=DB_HOST,
            user=DB_USER,
            passwd=DB_PASSWORD,
            database="sixecho")
    except mysql.connector.Error as err:
        if err.errno == errorcode.ER_ACCESS_DENIED_ERROR:
            print('Invalid credential. Unable to access database.')
        elif err.errno == errorcode.ER_BAD_DB_ERROR:
            print('Database does not exists')
        else:
            print('Failed to connect to database')
    mycursor = mydb.cursor()

    sql = "SELECT API_Secret FROM partners WHERE API_Key = '" + api_key + "'"
    mycursor.execute(sql)
    myresult = [x[0] for x in mycursor.fetchall()]
    # secret = myresult[0]
    # close cursor
    mycursor.close()
    # close connection
    mydb.close()
    return myresult  # not using

def validate_params(meta_books):
    language = meta_books["language"]
    if pycountry.languages.get(alpha_2=language) is None:
        raise Exception("ISO 639-1 Language is invalid.")
    country_of_origin = meta_books["country_of_origin"]
    if pycountry.countries.get(alpha_3=country_of_origin) is None:
        raise Exception("ISO 3166-1 Country is invalid.")
    publisher_id = meta_books["publisher_id"]
    check_publisher(publisher_id)
    category_id = meta_books["category_id"]
    check_category(category_id)


def check_publisher(publisher_id):
    session = Session()
    publisher = session.query(Publisher).filter_by(id=publisher_id).first()
    if publisher is None:
        raise Exception("Publisher ID is not exist")
    session.close()


def check_category(category_id):
    session = Session()
    category = session.query(Category).filter_by(id=category_id).first()
    if category is None:
        raise Exception("Category ID is not exist")
    session.close()

def sorted_toString(unsorted_dict):
    s = ''
    for key, value in sorted(unsorted_dict.items()):
        s = s + str(key) + str(value)
    return s
    # sortednames = sorted(unsorted_dict.keys(), key=lambda x: x.lower())
    # print(sortednames)
    # sorted_dict = {}
    # for i in sortednames:
    #     sorted_dict[i] = str(unsorted_dict[i])
    # return str(sorted_dict)


def create_sha256_signature(secret, message):
    secret = str(secret)
    message = str(message)
    secret_byte = str(secret).encode('utf-8')
    message_byte = str(message).encode('utf-8')
    signature = hmac.new(secret_byte, message_byte, hashlib.sha256).hexdigest()
    return signature

def lambda_handler(event, context):
    api_key = event["context"]["api-key"]
    sign = event["params"]["header"]["x-api-sign"]
    api_secret = get_API_Secret(api_key)[0]
    body = event["body-json"]
    for field in ["meta_books"]:
        if field not in body.keys():
            raise Exception("require %s argument." % field)
    unsorted_dict = body["meta_books"]
    sorted_dict = sorted_toString(unsorted_dict)
    signature = create_sha256_signature(str(api_secret), str(sorted_dict))
    if(signature == sign):
        host, redis_url, port = os.environ["REDIS_URL"].split(":")
        redis_url = redis_url.replace("//", "")
        print({'host': redis_url, 'port': port})

        lsh = MinHashLSH(
            storage_config={
                'type': 'redis',
                'redis': {
                    'host': redis_url,
                    'port': port
                },
                'basename': b'digital_checker',
            })
        uid = uuid.uuid4().hex
        body = event["body-json"]
        print(body)
        api_key_id = event["context"]["api-key-id"]
        try:
            for field in ["digest", "sha256", "size_file", "meta_books"]:
                if field not in body.keys():
                    raise Exception("require %s argument." % field)
            digest_str = body["digest"]
            sha256 = body["sha256"]
            size_file = body["size_file"]
            meta_books = body["meta_books"]
            validate_params(meta_books)
        except Exception as e:
            return {"statusCode": 200, "body": json.dumps({"message": e.message})}
        m1 = convert_str_to_minhash(digest_str)
        result = lsh.query(m1)
        if len(result) > 0:
            return {
                "statusCode": 200,
                "body": json.dumps({
                    "message": "Duplicate",
                }),
            }
        else:
            # call to another lambda
            meta_books = body["meta_books"]
            msg = {
                "name" : "new-book",
                "body" : {
                    "id" : uid,
                    "title" : meta_books["title"],
                    "author" : meta_books["author"],
                    "origin" : meta_books["country_of_origin"],
                    "lang" : meta_books["language"],
                    "paperback" : meta_books["paperback"],
                    "publisher_id" : meta_books["publisher_id"],
                    "publish_date" : meta_books["publish_date"]
                }
            }

            print(os.environ['CONTRACT_CLIENT_FUNCTION'])

            invoke_response = lambda_client.invoke(
                FunctionName=os.environ['CONTRACT_CLIENT_FUNCTION'],
                InvocationType="Event",
                Payload=json.dumps(msg))

            print(invoke_response)

            insert_mysql(api_key_id, uid, digest_str, sha256, size_file,
                        meta_books)
            lsh.insert(key=uid, minhash=m1)
            return {
                "statusCode": 200,
                "body": json.dumps({
                    "message": "Ok",
                }),
            }
    else: return {"statusCode": 403, "body": json.dumps({"message": "Signature Does Not Match"})}