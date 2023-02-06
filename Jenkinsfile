pipeline {
    agent any
    options {
	timeout(time: 1, unit: 'HOURS')
	timestamp()
	}
    stages {
        stage('Build') {
            steps {
                sh 'docker-compose build'
	          }
        }
        stage('Deploy') {
            steps {
               sh 'docker-compose up -d'
           }
        }
    }
}

