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
        withDockerRegistry(credentialsId: 'c77fa9f5-ab60-4075-8870-bbb1c1b26a3f') {
            sh label: '', script 'docker build -t ${REPO}/${IMAGE} .'
        }
      }
    }

    stage('Upload Img') {
    steps {
        withDockerRegistry(credentialsId: 'c77fa9f5-ab60-4075-8870-bbb1c1b26a3f') {
            sh label: '', script 'docker push ${REPO}/${IMAGE}'
        }
        sh label '', 'docker build -t ${REPO}/${IMAGE} .'
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