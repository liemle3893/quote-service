@Library("Demo Shared Library") _

pipeline {
	stages {
		stage('Build') {
			sh "make build"
		}
		stage('Test') {
			sh "mkdir -p build"
			sh "go test -v 2>&1 | go-junit-report > build/report.xml"
		}
		stage('Package') {
			docker.withRegistry('https://registry.hub.docker.com/', 'dockerhub') {
				docker.build('liemlhd/america-election-quote').push('latest')
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


