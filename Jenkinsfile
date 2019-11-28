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
						docker.build('saboteurkid/america-election-quote').push('latest')
					}
				}
			}
		}
		stage('Deploy') {
			steps {
				def userInput = input(
					id: 'userInput', message: "Deploy?", parameters: [
					[$class: 'BooleanParameterDefinition', defaultValue: false, description: '', name: 'Do you want to proceed?']
				])
				if(userInput) {
					script {
						sh "nomad stop quote-service"
						sh "nomad plan deployment/job.hcl"
						sh "nomad run deployment/job.hcl"
					}
				} else {
					echo "Deployment aborted"
				}
			}
		}
	}
	post {
        always {
			script {
				msteams()
			}
        }
    }	
}



