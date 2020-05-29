Safety Helmet

What Does The Program Do:
- Training Dataset Collected from Google and Baidu, Referenced from: https://github.com/njvisionpower/Safety-Helmet-Wearing-Dataset
- Trained using AWS Sagemaker Resnet-50 network using the pre-trained model, validation mAP: 63
- Inference result for two classes: ['helmet', 'person'] for person with helmet, and person without helmet
- RTSP feed transfer to Kinesis Stream, using the model trained to detect helmet/no_helmet, send to kinesis datastream, using Lambda to save information to S3
- Using web and lambda API to display information from S3

Future Works:
- Fine tune the current model with real security camera on site
- integrate the APIs to the video management systems
- Refine the cloud formation for deployment scalability and robustness
- Implement detection classes for safety cloth and safety shoes