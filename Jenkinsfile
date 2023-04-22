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
        withDockerRegistry(credentialsId: 'c7df7343-cf32-4500-8161-dd42a0a26451', url: 'https://index.docker.io/v1/') {
            sh label: '', script: 'docker build -t ${REPO}/${IMAGE} .'
        }
      }
    }

    stage('Upload Img') {
    steps {
        withDockerRegistry(credentialsId: 'c7df7343-cf32-4500-8161-dd42a0a26451', url: 'https://index.docker.io/v1/') {
            sh label: '', script: 'docker push ${REPO}/${IMAGE}'
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