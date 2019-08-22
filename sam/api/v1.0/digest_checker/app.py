import mysql.connector
from mysql.connector import errorcode
import boto3
import hashlib
import hmac
import json
import os
from sqlalchemy import create_engine
from sqlalchemy.orm import sessionmaker

import numpy as np
import text

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
    return myresult  # not using

 
def validate_image(event):
    print(event)
     
def lambda_handler(event, context):
    api_key = event["context"]["api-key"]
    sign = event["params"]["header"]["x-api-sign"]
    api_secret = get_api_secret(api_key)[0]
    body = event["body-json"]
    for field in ["meta_media"]:
        if field not in body.keys():
            raise Exception("require %s argument." % field)
    unsorted_dict = body["meta_media"]
    sorted_dict = sorted_to_string(unsorted_dict)
    signature = create_sha256_signature(str(api_secret), str(sorted_dict))
    if (signature != sign):
        return {
            "statusCode": 403,
            "body": json.dumps({"message": "Signature Does Not Match"})
        }
    if body["type"] == "TEXT":
        text.validate_text(Session,event)
    elif body["type"] == "IMAGE":
        validate_image(event)
