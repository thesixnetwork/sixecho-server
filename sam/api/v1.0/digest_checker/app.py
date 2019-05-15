from datasketch import MinHash, MinHashLSH, LeanMinHash
import json
import os


def lambda_handler(event, context):
    print(event)
    print("---------------")
    print(context)
    print("VVVVVVVVVVVVVVV")
    redis_url, port = os.environ["REDIS_URL"].split(":")
    lsh = MinHashLSH(
        storage_config={
            'type': 'redis',
            'redis': {'host': redis_url, 'port': port}
        })
    digest = event["digest"]
    md5 = event["md5"]
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
        lsh.insert(m1)
        return{
            "statusCode": 200,
            "body": json.dumps({
                "message": "hello world",
            }),
        }
