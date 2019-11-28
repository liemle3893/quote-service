@Library("Demo Shared Library") _

pipeline {
	agent any
	environment {
		MAJOR_VERSION = "1"
		MINOR_VERSION = "0"
		PATCH_VERSION = "${currentBuild.number}"
	}
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
						image = docker.build('saboteurkid/america-election-quote')
						image.push("${MAJOR_VERSION}.${MINOR_VERSION}.${PATCH_VERSION}")
						image.push('latest')
					}
				}
			}
		}
		stage('Deploy') {
			steps {
				script {
					// Verifu deployment
					def userInput = input(
						id: 'userInput', message: "Deploy?", parameters: [
						[$class: 'BooleanParameterDefinition', defaultValue: false, description: '', name: 'Do you want to proceed?']
					])
					if(userInput) {
						script {
							// Deploykent
							sh "nomad stop quote-service || true"
							sh "nomad plan deployment/job.hcl || true"
							sh "nomad run deployment/job.hcl"
						}
					} else {
						echo "Deployment aborted"
					}
				}
			}
		}
	}
	post {
        always {
			script {
				junit "build/*.xml"
				msteams()
			}
        }
    }	
}



