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
        withDockerRegistry(credentialsId: '10a409e9-eb12-4f6c-833b-4ad8c17abc3d', url: 'https://index.docker.io/v1/') {
            sh label: '', script: 'sudo docker build -t ${REPO}/${IMAGE} .'
        }
      }
    }

    stage('Upload Img') {
    steps {
        withDockerRegistry(credentialsId: '10a409e9-eb12-4f6c-833b-4ad8c17abc3d', url: 'https://index.docker.io/v1/') {
            sh label: '', script: 'sudo docker push ${REPO}/${IMAGE}'
        }
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