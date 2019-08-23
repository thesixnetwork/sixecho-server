import uuid
import json
from datetime import datetime
from sixecho_model import Category, DigitalContent, HashImage


def check_hash_image(Session, id):
    session = Session()
    publisher = session.query(HashImage).filter_by(id=id).first()
    if publisher is not None:
        raise Exception("Duplicate Image ID "+ id)
    session.close()


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
     
    average_hash, phash, dhash, whash = body["digest"].split(",")
    ahash = HashImage(id=average_hash,hash_type="ahash") 
    phash = HashImage(id=phash,hash_type="phash") 
    dhash = HashImage(id=dhash,hash_type="dhash") 
    whash = HashImage(id=whash,hash_type="whash") 
    session.add(ahash)
    session.add(phash)
    session.add(dhash)
    session.add(whash)
    session.commit()

def validate(Session, event):
    uid = uuid.uuid4().hex
    body = event["body-json"]
    api_key_id = event["context"]["api-key-id"]
    average_hash, phash, dhash, whash = body["digest"].split(",")
    try:
      check_hash_image(Session,average_hash)  
      check_hash_image(Session,phash)  
      check_hash_image(Session,dhash)  
      check_hash_image(Session,whash)  
      insert_mysql(Session, api_key_id, uid, body)
      return {
            "statusCode": 200,
            "body": json.dumps({
                "message": "Ok",
                "id": uid,
      })}
    except Exception as e:
        print("Error " + str(e))
        return {"statusCode": 200, "body": json.dumps({"message": str(e)})}
