pipeline {
    agent any
    stages {
        stage('Build') {
            steps {
                sh 'docker-compose build'
		options {
			timeout(time: 1, unit: 'HOURS')
			timestamp()			
	}
            }
        }
        stage('Deploy') {
            steps {
               sh 'docker-compose up -d'
           }
        }
    }
}

