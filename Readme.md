# Safety Helmet

## What Does The Program Do:
- Training Dataset Collected from Google and Baidu
- Trained using AWS Sagemaker Resnet-50 network using ImageNet pre-trained model, validation mAP: 62.82
- Inference result for two classes: ['helmet', 'person'] for person with helmet, and person without helmet
- RTSP feed transfer to Kinesis Stream, using the model trained to detect helmet/no_helmet, send to kinesis datastream, using Lambda to save information to S3
- Using web and lambda API to display information from S3

## Future Works:
- Fine tune the current model with real security camera on site
- integrate the APIs to the video management systems
- Refine the cloud formation for deployment scalability and robustness
- Implement detection classes for safety cloth and safety shoes

## Demo:
<p align="center"> 
<img src="https://github.com/NTTDataCloud/SafetyHelmetSagemaker/raw/master/demo1.png" width = 50% height = 50%>
</p>  
<p align="center"> 
<img src="https://github.com/NTTDataCloud/SafetyHelmetSagemaker/raw/master/demo2.png" width = 50% height = 50%>
</p>  

# Demo Website: http://happyhelmetpostprocalb-518099815.us-east-1.elb.amazonaws.com/
The website demonstrates video frames which has persons without wearing helmet.