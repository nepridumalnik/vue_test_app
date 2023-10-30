<template>
    <div>
        <h1>{{ name }}</h1>
        <input type="number" :value="counter" @input="validateAndSet(Number($event.target.value))">
    </div>
</template>

<script>
export default {
    name: 'CountSmtn',
    props: {
        name: String,
        min: Number,
        max: Number,
    },
    created() {
        if (!(this.min < this.max)) {
            throw new Error('Min value ' + this.min + ' nor lower then max value ' + this.max)
        }
    },
    data() {
        return {
            counter: this.min,
        }
    },
    methods: {
        validateAndSet(value) {
            this.counter = value

            if (this.counter < this.min) {
                this.counter = this.min
            }
            else if (value > this.max) {
                this.counter = this.max
            }
        }
    }
}
</script>