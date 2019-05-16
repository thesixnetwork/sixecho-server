from datasketch import MinHash, MinHashLSH, LeanMinHash
import numpy as np
import json
import os
import uuid
import mysql.connector
import boto3


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


def insert_mysql(api_key_id=None, book_id=None, digest=None, sha256=None):
    """Insert digital content to mysql for keeping information about the book.
    Args:
        api_key_id(string)  - Required  : api key from api gateway
        book_id(string)     - Required  : unique id we generate by time
        digest(string)      - Required  : digest from client
        sha256(string)      - Rquireed  : sha256 generate from client
    """
    mydb = mysql.connector.connect(
        host=DB_HOST,
        user=DB_USER,
        passwd=DB_PASSWORD
    )
    mycursor = mydb.cursor()

    sql = "INSERT INTO digital_contents (api_key_id, book_id, digest, sha256) VALUES (%s, %s,%s,%s)"
    val = (api_key_id, book_id, digest, sha256)
    mycursor.execute(sql, val)
    mydb.commit()
    mycursor.close()


def lambda_handler(event, context):
    print(event)
    print("----------------------------------")
    host, redis_url, port = os.environ["REDIS_URL"].split(":")
    redis_url = redis_url.replace("//", "")
    print({'host': redis_url, 'port': port})
    lsh = MinHashLSH(
        storage_config={
            'type': 'redis',
            'redis': {'host': redis_url, 'port': port},
            'basename': b'digital_checker',
        })
    uid = uuid.uuid4().hex
    try:
        digest_str = event["digest"]
        sha256 = event["sha256"]
    except:
        return {
            "statusCode": 200,
            "body": json.dumps({
                "message": "error args"
            })
        }
    m1 = convert_str_to_minhash(digest_str)
    result = lsh.query(m1)
    if len(result) > 0:
        return{
            "statusCode": 200,
            "body": json.dumps({
                "message": "Duplicate",
            }),
        }
    else:
        insert_mysql("xxx", uid, digest_str, sha256)
        lsh.insert(key=uid, minhash=m1)
        return{
            "statusCode": 200,
            "body": json.dumps({
                "message": "Ok",
            }),
        }
