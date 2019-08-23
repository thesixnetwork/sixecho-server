import json
import os
import uuid
from datetime import datetime

import boto3
import numpy as np
import pycountry
from datasketch import MinHash, MinHashLSH

from sixecho_model import Category, DigitalContent, Publisher

lambda_client = boto3.client('lambda')


def validate_params(Session, meta_books):
    print("Validate Param")
    language = meta_books["language"]
    if pycountry.languages.get(alpha_2=language) is None:
        raise Exception("ISO 639-1 Language is invalid.")
    country_of_origin = meta_books["country_of_origin"]
    if pycountry.countries.get(alpha_3=country_of_origin) is None:
        raise Exception("ISO 3166-1 Country is invalid.")
    #  publisher_id = meta_books["publisher_id"]
    #  check_publisher(Session, publisher_id)
    #  category_id = meta_books["category_id"]
    #  check_category(Session, category_id)


#  def check_publisher(Session, publisher_id):
#  print("Check Publisher")
#  session = Session()
#  publisher = session.query(Publisher).filter_by(id=publisher_id).first()
#  if publisher is None:
#  raise Exception("Publisher ID is not exist")
#  session.close()

#  def check_category(Session, category_id):
#  session = Session()
#  category = session.query(Category).filter_by(id=category_id).first()
#  if category is None:
#  raise Exception("Category ID is not exist")
#  session.close()


def convert_str_to_minhash(digest):
    """Convert string that is including 128 numbers which to have a comma as middle between that numbers.
    Ex. 13241234,213242134,22342234,23423423,...,21341234 (128 numbers.)
    """
    data_array = np.array(digest.split(","), dtype=np.uint64)
    m1 = MinHash(hashvalues=data_array)
    return m1


def send_to_chain(uid, body):
    meta_books = body["meta_media"]
    digest_str = body["digest"]
    digest_str_array = digest_str.split(',')
    digest_int_array = list(map(int, digest_str_array))
    msg = {
        "name": "new-book-and-digest",
        "body": {
            "id": uid,
            "title": meta_books["title"],
            "author": meta_books["author"],
            "origin": meta_books["country_of_origin"],
            "lang": meta_books["language"],
            "paperback": meta_books["paperback"],
            "publisher_id": meta_books["publisher_id"],
            "publish_date": meta_books["publish_date"],
            "digest": digest_int_array
        }
    }

    print(msg)

    print(os.environ['CONTRACT_CLIENT_FUNCTION'])

    invoke_response = lambda_client.invoke(
        FunctionName=os.environ['CONTRACT_CLIENT_FUNCTION'],
        InvocationType="Event",
        Payload=json.dumps(msg))


def insert_mysql(Session, api_key_id=None, uid=None, body=None):
    session = Session()
    digest_str = body["digest"]
    sha256 = body["sha256"]
    size_file = body["size_file"]
    content_type = body["type"]
    meta_books = body["meta_media"]

    json_meta_books = json.dumps(meta_books)
    category_id = meta_books["category_id"]
    publisher_id = meta_books["publisher_id"]
    title = meta_books["title"]
    author = meta_books["author"]
    # country_of_origin = meta_books["country_of_origin"]
    # language = meta_books["language"]
    # paperback = meta_books["paperback"]
    # publish_date = meta_books["publish_date"]
    digital_content_id = meta_books.get("digital_content_id")
    digital_content = DigitalContent(api_key_id=api_key_id,
                                     id=uid,
                                     category_id=category_id,
                                     publisher_id=publisher_id,
                                     digital_content_id=digital_content_id,
                                     title=title,
                                     digest=digest_str,
                                     sha256=sha256,
                                     size_file=size_file,
                                     content_type=content_type,
                                     author=author,
                                     meta_media=json_meta_books,
                                     created_at=datetime.now(),
                                     updated_at=datetime.now())
    session.add(digital_content)
    session.commit()


def validate(Session, event):
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
        digest_str = body["digest"]
        meta_books = body["meta_media"]
        validate_params(Session, meta_books)
    except Exception as e:
        print("Error " + str(e))
        return {"statusCode": 200, "body": json.dumps({"message": str(e)})}
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
        insert_mysql(Session, api_key_id, uid, body)
        lsh.insert(key=uid, minhash=m1)
        return {
            "statusCode": 200,
            "body": json.dumps({
                "message": "Ok",
                "id": uid,
            }),
        }
