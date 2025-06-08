pipeline {
    agent any
    environment {
        REPO = 'https://github.com/ev-smoke/telebot-go'
        BRANCH = 'main'
    }
    
    parameters {
        choice(
            name: 'OS',
            choices: ['linux', 'darwin', 'windows'],
            description: 'Target operating system'
        )
        choice(
            name: 'ARCH',
            choices: ['amd64', 'arm64'],
            description: 'Target architecture'
        )
        booleanParam(
            name: 'SKIP_TESTS',
            defaultValue: false,
            description: 'Skip running tests'
        )
        booleanParam(
            name: 'SKIP_LINT',
            defaultValue: false,
            description: 'Skip running linter'
        )
    }

    stages {
        stage ("clone") {
            steps {
                echo '>>> Clone repository'
                git branch: "${BRANCH}", url: "${REPO}"
            }
        }

        stage ("build") {
            steps {
                echo '>>> Build packet'
                sh 'make build'
            }
        }

        stage ("image") {
            steps {
                echo '>>> Build image'
                sh 'make image'
            }
        }

        stage ("push") {
            steps {
                echo '>>> push image'
                script {
                    docker.withRegistry( '', 'dockerhub' ) {
                        sh 'make push'
                    }
                }
            }
        }
    }
}
