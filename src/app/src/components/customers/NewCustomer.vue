<template>
    <form @submit="submitForm">
        <label for="">Name</label><br>
        <input type="text" v-model="customer.name"><br>
        <label for="">Email</label><br>
        <input type="email" v-model="customer.email"><br>
        <label for="">Phone</label><br>
        <input type="tel" v-model="customer.phone"><br>
        <label for="">Street</label><br>
        <input type="text" v-model="customer.street"><br>
        <label for="">Neighborhood</label><br>
        <input type="text" v-model="customer.neighborhood"><br>
        <label for="">Number</label><br>
        <input type="number" v-model="customer.number"><br>
        <label for="">zip code</label>  <br>
        <input type="text" v-model="customer.zipcode"><br>
        <button type="submit">Save</button>
    </form>
</template>

<script>
import axios from 'axios'

export default {
    data (){
        return {
            customer: {
                name: '',
                email: '',
                phone: '',
                street: '',
                number: 0,
                neighborhood: '',
                zipcode: ''
            }
        }
    },
    methods: {
        submitForm: function(e) {
            e.preventDefault()
            alert(this.customer)
            axios.post("http://localhost:8081/customers/", this.customer)
            .then(resp => window.location.href = "#/lista")
        }
    }
}
</script>

<style>

</style>
