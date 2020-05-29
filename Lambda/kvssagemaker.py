from __future__ import print_function
import base64
import json
import boto3
import os
import datetime
import time
from botocore.exceptions import ClientError

bucket='<<YOUR BUCKET>>'

#Lambda function is written based on output from an Amazon SageMaker example: 
#https://github.com/awslabs/amazon-sagemaker-examples/blob/master/introduction_to_amazon_algorithms/object_detection_pascalvoc_coco/object_detection_image_json_format.ipynb
object_categories = ['person', 'bicycle', 'car',  'motorbike', 'aeroplane', 'bus', 'train', 'truck', 'boat', 
                     'traffic light', 'fire hydrant', 'stop sign', 'parking meter', 'bench', 'bird', 'cat', 'dog',
                     'horse', 'sheep', 'cow', 'elephant', 'bear', 'zebra', 'giraffe', 'backpack', 'umbrella', 'handbag',
                     'tie', 'suitcase', 'frisbee', 'skis', 'snowboard', 'sports ball', 'kite', 'baseball bat',
                     'baseball glove', 'skateboard', 'surfboard', 'tennis racket', 'bottle', 'wine glass', 'cup',
                     'fork', 'knife', 'spoon', 'bowl', 'banana', 'apple', 'sandwich', 'orange', 'broccoli', 'carrot',
                     'hot dog', 'pizza', 'donut', 'cake', 'chair', 'sofa', 'pottedplant', 'bed', 'diningtable',
                     'toilet', 'tvmonitor', 'laptop', 'mouse', 'remote', 'keyboard', 'cell phone', 'microwave', 'oven',
                     'toaster', 'sink', 'refrigerator', 'book', 'clock', 'vase', 'scissors', 'teddy bear', 'hair drier',
                     'toothbrush']

def lambda_handler(event, context):
  for record in event['Records']:
    payload = base64.b64decode(record['kinesis']['data'])
    #Get Json format of Kinesis Data Stream Output
    result = json.loads(payload)
    #Get FragmentMetaData
    fragment = result['fragmentMetaData']
    
    # Extract Fragment ID and Timestamp
    frag_id = fragment[17:-1].split(",")[0].split("=")[1]
    srv_ts = datetime.datetime.fromtimestamp(float(fragment[17:-1].split(",")[1].split("=")[1])/1000)
    srv_ts1 = srv_ts.strftime("%A, %d %B %Y %H:%M:%S")
    
    #Get FrameMetaData
    frame = result['frameMetaData']
    #Get StreamName
    streamName = result['streamName']
   
    #Get SageMaker response in Json format
    sageMakerOutput = json.loads(base64.b64decode(result['sageMakerOutput']))
    #Print 5 detected object with highest probability
    for i in range(5):
      print("detected object: " + object_categories[int(sageMakerOutput['prediction'][i][0])] + ", with probability: " + str(sageMakerOutput['prediction'][i][1]))
    
    detections={}
    detections['StreamName']=streamName
    detections['fragmentMetaData']=fragment
    detections['frameMetaData']=frame
    detections['sageMakerOutput']=sageMakerOutput

    #Get KVS fragment and write .webm file and detection details to S3
    s3 = boto3.client('s3')
    kv = boto3.client('kinesisvideo')
    get_ep = kv.get_data_endpoint(StreamName=streamName, APIName='GET_MEDIA_FOR_FRAGMENT_LIST')
    kvam_ep = get_ep['DataEndpoint']
    kvam = boto3.client('kinesis-video-archived-media', endpoint_url=kvam_ep)
    getmedia = kvam.get_media_for_fragment_list(
                            StreamName=streamName,
                            Fragments=[frag_id])
    base_key=streamName+"_"+time.strftime("%Y%m%d-%H%M%S")
    webm_key=base_key+'.webm'
    text_key=base_key+'.txt'
    s3.put_object(Bucket=bucket, Key=webm_key, Body=getmedia['Payload'].read())
    s3.put_object(Bucket=bucket, Key=text_key, Body=json.dumps(detections))
    print("Detection details and fragment stored in the S3 bucket "+bucket+" with object names : "+webm_key+" & "+text_key)
  return 'Successfully processed {} records.'.format(len(event['Records']))
