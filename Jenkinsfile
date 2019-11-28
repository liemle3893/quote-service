@Library("Demo Shared Library") _

pipeline {
	agent any
	stages {
		stage('Test') {
			steps {
				sh "mkdir -p build"
				sh "go test -v 2>&1 | go-junit-report > build/report.xml"
			}
		}
		stage('Package') {
			steps {
				script {
					docker.withRegistry('https://registry.hub.docker.com/', 'dockerhub') {
						docker.build('liemlhd/america-election-quote').push('latest')
					}
				}
			}
		}
	}
	post {
        always {
            junit 'build/*.xml'
			script {
				msteams()
			}
        }
    }	
}



