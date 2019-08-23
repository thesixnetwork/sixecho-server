import hashlib
import hmac
import json
import os

import boto3
import mysql.connector
import numpy as np
from mysql.connector import errorcode
from sqlalchemy import create_engine
from sqlalchemy.orm import sessionmaker

import image
import text
from sixecho_model import Category, DigitalContent, Publisher

lambda_client = boto3.client('lambda')

ssm = boto3.client('ssm')

parameter = ssm.get_parameter(Name="SIXECHO_HOST_DB")
parameter2 = ssm.get_parameter(Name="SIXECHO_USER_DB")
parameter3 = ssm.get_parameter(Name="SIXECHO_PASSWORD_DB")

DB_HOST = parameter["Parameter"]["Value"]
DB_USER = parameter2["Parameter"]["Value"]
DB_PASSWORD = parameter3["Parameter"]["Value"]

URL_ENGINE = "mysql+mysqlconnector://%s:%s@%s/sixecho" % (DB_USER, DB_PASSWORD,
                                                          DB_HOST)
ENGINE = create_engine(URL_ENGINE, echo=True)
Session = sessionmaker(bind=ENGINE)


def sorted_to_string(unsorted_dict):
    s = ''
    for key, value in sorted(unsorted_dict.items()):
        s = s + str(key) + str(value)
    return s


def create_sha256_signature(secret, message):
    secret = str(secret)
    message = str(message)
    secret_byte = str(secret).encode('utf-8')
    message_byte = str(message).encode('utf-8')
    signature = hmac.new(secret_byte, message_byte, hashlib.sha256).hexdigest()
    return signature


def get_api_secret(api_key):
    try:
        mydb = mysql.connector.connect(host=DB_HOST,
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
    mycursor.close()
    mydb.close()
    print("Result ---")
    print(myresult)
    print("@@@@@@@@@@")
    return myresult  # not using


def validate_image(event):
    print(event)


def check_publisher(publisher_id):
    print("Check Publisher")
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


def lambda_handler(event, context):
    api_key = event["context"]["api-key"]
    sign = event["params"]["header"]["x-api-sign"]
    api_secret = get_api_secret(api_key)[0]
    body = event["body-json"]
    try:
        for field in [
                "digest", "sha256", "size_file", "meta_media", "category_id",
                "publisher_id"
        ]:
            if field not in body.keys():
                raise Exception("require %s argument." % field)
        meta_media = body["meta_media"]
        check_publisher(meta_media["publisher_id"])
        check_category(meta_media["category_id"])
    except Exception as e:
        print("Error " + str(e))
        return {"statusCode": 200, "body": json.dumps({"message": str(e)})}

    unsorted_dict = body["meta_media"]
    sorted_dict = sorted_to_string(unsorted_dict)
    signature = create_sha256_signature(str(api_secret), str(sorted_dict))
    if (signature != sign):
        return {
            "statusCode": 403,
            "body": json.dumps({"message": "Signature Does Not Match"})
        }

    if body["type"] == "TEXT":
        return text.validate(Session, event)
    elif body["type"] == "IMAGE":
        return image.validate(Session, event)
