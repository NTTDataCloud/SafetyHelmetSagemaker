Safety Helmet

What Does The Program Do:
- Training Dataset Collected from Google and Baidu, Referenced from: https://github.com/njvisionpower/Safety-Helmet-Wearing-Dataset
- Trained using AWS Sagemaker Resnet-50 network using the pre-trained model.
- Training dataset in VOC like data format, converted to Sagemaker JSON format using Notebook
- Training valuation: mAP: 63
- Inference result for two classes: ['helmet', 'person'] for person with helmet, and person without helmet
