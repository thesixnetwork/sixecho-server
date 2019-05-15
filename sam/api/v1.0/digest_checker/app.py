from datasketch import MinHash, MinHashLSH, LeanMinHash
import json
import os
import uuid


def lambda_handler(event, context):
    redis_url, port = os.environ["REDIS_URL"].split(":")
    lsh = MinHashLSH(
        storage_config={
            'type': 'redis',
            'redis': {'host': redis_url, 'port': port}
        })
    uid = uuid.uuid4().hex
    try:
        digest = event["digest"]
        md5 = event["md5"]
    except:
        return {
            "statusCode": 200,
            "body": json.dumps({
                "message": "error args"
            })
        }
    m1 = LeanMinHash.deserialize(digest)
    result = lsh.query(m1)
    if len(result) > 0:
        return{
            "statusCode": 200,
            "body": json.dumps({
                "message": "hello world",
            }),
        }
    else:
        lsh.insert(key=uid, minhash=m1)
        return{
            "statusCode": 200,
            "body": json.dumps({
                "message": "hello world",
            }),
        }
