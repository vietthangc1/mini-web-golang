pipeline {
  agent any
  stages {
    stage('Start') {
      steps {
        echo 'Start Jenkins!'
      }
    }

    stage('Build Img') {
      steps {
            sh 'sudo docker build -t ${REPO}/${IMAGE} .'
        }
    }

    stage('Upload Img') {
    steps {
            sh 'sudo docker push ${REPO}/${IMAGE}'
      }
    }

    stage('End') {
      steps {
        echo 'Finish Jenkins'
      }
    }

  }
  environment {
    IMAGE = 'mini-golang-project'
    REPO = 'vietthangc1'
  }
}