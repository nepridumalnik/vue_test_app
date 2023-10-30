<template>
    <div class="root">
        <label>
            <span>Почта</span>
            <input type="email" :value="email" @input="emailInput($event.target.value)">
        </label>
        <label>
            <span>Пароль</span>
            <input type="password" :value="password" @input="passwordInput($event.target.value)">
        </label>
    </div>
</template>

<script>
export default {
    name: 'MyComponent',
    data() {
        return {
            email: 'some@mail.org',
            password: 'хуйхуйхуй',
        }
    },
    methods: {
        emailInput(email) {
            this.email = email
            console.log(this.email)
        },
        passwordInput(password) {
            this.password = password
            console.log(this.password)
        },
    }
}
</script>
