pipeline {

  agent none

  environment {
    MESSAGE_TEST = "Running Jenkins"
  }

  stages {
    stage("Test") {
      steps {
        sh "echo ${MESSAGE_TEST}"
      }
    }

    stage("build") {
      agent { node {label 'main'}}
      environment {
      }
      steps {
      }
    }
  }
}