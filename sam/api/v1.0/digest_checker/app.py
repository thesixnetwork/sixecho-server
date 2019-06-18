import json
import os
import uuid
from datetime import datetime

import boto3
import mysql.connector
import numpy as np
from datasketch import LeanMinHash, MinHash, MinHashLSH

ssm = boto3.client('ssm')

parameter = ssm.get_parameter(Name="SIXECHO_HOST_DB")
parameter2 = ssm.get_parameter(Name="SIXECHO_USER_DB")
parameter3 = ssm.get_parameter(Name="SIXECHO_PASSWORD_DB")

DB_HOST = parameter["Parameter"]["Value"]
DB_USER = parameter2["Parameter"]["Value"]
DB_PASSWORD = parameter3["Parameter"]["Value"]


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


def lambda_handler(event, context):
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
    api_key_id = event["context"]["api-key-id"]
    try:
        digest_str = body["digest"]
        sha256 = body["sha256"]
        size_file = body["size_file"]
        meta_books = body["meta_books"]
    except:
        return {
            "statusCode": 200,
            "body": json.dumps({"message": "error args"})
        }
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
        insert_mysql(api_key_id, uid, digest_str, sha256, size_file,
                     meta_books)
        lsh.insert(key=uid, minhash=m1)
        return {
            "statusCode": 200,
            "body": json.dumps({
                "message": "Ok",
            }),
        }
