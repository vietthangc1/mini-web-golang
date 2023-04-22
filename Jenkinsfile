pipeline {
  agent none
  stages {
    stage("Test") {
    agent {
            docker { image 'golang:1.18-alpinee' }
        }
      steps {
        sh "go --version"
      }
    }
  }
}