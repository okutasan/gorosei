pipeline {
  agent {
    docker {
      image 'openjdk:8-alpine'
      args '-v /tmp:/tmp'
      reuseNode true
    }
  }
  options {
        copyArtifactPermission('../shenlong/*');
    }
  stages {
    stage('Build') {
      steps {
        echo 'Build Kohmon'
        sh 'rm -rf builds'
        sh 'mkdir builds'
        sh 'cp *.go builds'
        sh 'ls builds'
      }
    }

  }
  post {
        always {
            archiveArtifacts artifacts: 'builds/*.go', fingerprint: true
        }
    }
}

