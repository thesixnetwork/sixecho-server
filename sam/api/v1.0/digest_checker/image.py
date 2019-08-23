import uuid

from sixecho_model import Category, DigitalContent, Publisher


def check_digital_content(Session, imghash):
    print("")


def validate(Session, event):
    uid = uuid.uuid4().hex
    body = event["body-json"]
    api_key_id = event["context"]["api-key-id"]
    average_hash, phash, dhash, whash = body["digest"].split(",")
