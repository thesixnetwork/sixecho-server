from datasketch import MinHash, MinHashLSH, LeanMinHash
import numpy as np
import json
import os
import uuid


def convert_str_to_minhash(digest):
    data_array = np.array(digest.split(","), dtype=np.uint64)
    m1 = MinHash(hashvalues=data_array)
    return m1


def lambda_handler(event, context):
    host, redis_url, port = os.environ["REDIS_URL"].split(":")
    lsh = MinHashLSH(
        storage_config={
            'type': 'redis',
            'redis': {'host': redis_url, 'port': port}
        })
    uid = uuid.uuid4().hex
    try:
        digest_str = event["digest"]
        md5 = event["sha256"]
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
