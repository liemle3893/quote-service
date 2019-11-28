@Library("Demo Shared Library") _

pipeline {
	agent any
	stages {
		stage('Build') {
			agent {
				docker {
					image 'golang:1.13.0-stretch'
				}
			}
			steps {
				sh "mkdir -p build"
				sh "go build -mod=vendor"
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



