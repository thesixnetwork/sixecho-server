import json
import uuid

import numpy as np
import pycountry
from datasketch import LeanMinHash, MinHash, MinHashLSH

from .sixecho_model.category import Category
from .sixecho_model.publisher import Publisher


def validate_params(Session, meta_books):
    language = meta_books["language"]
    if pycountry.languages.get(alpha_2=language) is None:
        raise Exception("ISO 639-1 Language is invalid.")
    country_of_origin = meta_books["country_of_origin"]
    if pycountry.countries.get(alpha_3=country_of_origin) is None:
        raise Exception("ISO 3166-1 Country is invalid.")
    publisher_id = meta_books["publisher_id"]
    check_publisher(Session, publisher_id)
    category_id = meta_books["category_id"]
    check_category(Session, category_id)


def check_publisher(Session, publisher_id):
    session = Session()
    publisher = session.query(Publisher).filter_by(id=publisher_id).first()
    if publisher is None:
        raise Exception("Publisher ID is not exist")
    session.close()


def check_category(Session, category_id):
    session = Session()
    category = session.query(Category).filter_by(id=category_id).first()
    if category is None:
        raise Exception("Category ID is not exist")
    session.close()


def convert_str_to_minhash(digest):
    """Convert string that is including 128 numbers which to have a comma as middle between that numbers.
    Ex. 13241234,213242134,22342234,23423423,...,21341234 (128 numbers.)
    """
    data_array = np.array(digest.split(","), dtype=np.uint64)
    m1 = MinHash(hashvalues=data_array)
    return m1


def validate_text(Session, event):
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
        for field in ["digest", "sha256", "size_file", "meta_media"]:
            if field not in body.keys():
                raise Exception("require %s argument." % field)
        digest_str = body["digest"]
        sha256 = body["sha256"]
        size_file = body["size_file"]
        meta_books = body["meta_media"]
        validate_params(Session, meta_books)
    except Exception as e:
        return {"statusCode": 200, "body": json.dumps({"message": str(e)})}
    m1 = convert_str_to_minhash(digest_str)
