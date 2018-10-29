var app = new Vue({
    el: '#quiz',
    data: {
        selectedAnswer: null,
        questions: [],
        currentQuestion: 0,
        stage: {
            intro: true,
            quiz: false,
            results: false
        },
        errors: {
            name: false,
            email: false
        },
        user: {
            name: null,
            email: null,
            difficulty: 'Easy'
        },
        answers: [],
        correct: 0,
        perc: null,
        rank: 0.0
    },
    methods: {
        processForm: function() {
            if (this.errors.email || this.errors.name) {
                return false
            }
            axios.post('/quiz', {
                name: this.user.name,
                email: this.user.email,
                difficulty: this.user.difficulty
            }).then((response) => {
                this.stage.intro = false
                this.stage.quiz = true
                this.questions = response.data.QuestionsPack
            }).catch((e) => {
                console.error(e)
            })
        },
        validateEmail: function() {
            const isValid = isValidEmail(this.user.email);
            this.errors.email = !isValid;
        },
        handleAnswer(e) {
            this.answers[this.currentQuestion] = e.answer;
            if ((this.currentQuestion + 1) === this.questions.length) {
                this.handleResults();
                this.stage.quiz = false;
                this.stage.results = true;
            } else {
                this.currentQuestion++;
            }
        },
        handleResults() {
            // Post results to the server
            axios.post('/results', {
                user: this.user,
                questions: this.questions,
                answers: this.answers
            }).then((response) => {
                this.correct = response.data.correct_answers
                this.perc = response.data.percentage
                this.rank = response.data.rank
            }).catch((e) => {
                console.error(e)
            })
        }
    }
});

function isValidEmail(email) {
    var re = /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
    return re.test(email);
}


Vue.component('question', {
    template: `
<div id="quiz">
  <strong>Question {{ questionNumber }}:</strong><br/>
  <strong>{{ question.question }} </strong>
<br />
<br />
<div v-for="a of question.answers">
  <input type="radio" :name="a" :value="a" v-model="answer">&nbsp;&nbsp;{{ a }}
</div>
<br />
<br />

<button class="button is-danger" @click="submitAnswer">Answer</button>
</div>
`,
    data() {
        return {
            answer: ''
        }
    },
    props: ['question', 'question-number'],
    methods: {
        submitAnswer: function() {
            this.$emit('answer', {
                answer: this.answer
            });
            this.answer = null;
        }
    }
});