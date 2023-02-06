pipeline {
    agent any
    options {
	timeout(time: 1, unit: 'HOURS')
	timestamps()
	ansiColor('xterm')
	}
    stages {
        stage('Build') {
            steps {
                sh 'docker-compose build'
		echo '\033[34mHello\033[0m \033[33mcolorful\033[0m \033[35mworld!\033[0m'
	          }
        }
        stage('Deploy') {
            steps {
               sh 'docker-compose up -d'
	       echo '\033[34mHello\033[0m \033[33mcolorful\033[0m \033[35mworld!\033[0m'
           }
        }
    }
}

