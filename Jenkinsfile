@Library("Demo Shared Library") _

pipeline {
	stages {
		stage('Build') {
			steps {
				sh "make build"
			}
		}
		stage('Test') {
			steps {
				sh "mkdir -p build"
				sh "go test -v 2>&1 | go-junit-report > build/report.xml"
			}
		}
		stage('Package') {
			steps {
				docker.withRegistry('https://registry.hub.docker.com/', 'dockerhub') {
					docker.build('liemlhd/america-election-quote').push('latest')
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



