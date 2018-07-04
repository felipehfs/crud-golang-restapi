<template>
    <table class="listCustomers">
        <caption>Lista de clientes</caption>
        <thead>
            <tr>
                <th>Name</th>
                <th>Email</th>
                <th>Phone</th>
                <th>Street</th>
                <th>Neighborhood</th>
                <th>Number</th>
                <th>ZipCode</th>
            </tr>
        </thead>
        <tbody>
            <tr v-for="customer in customers" :key="customer.id">
                <td>{{customer.name}}</td>
                <td>{{customer.email}}</td>
                <td>{{customer.phone}}</td>
                <td>{{customer.street}}</td>
                <td>{{customer.neighborhood}}</td>
                <td>{{customer.number}}</td>
                <td>{{customer.zipcode}}</td>
                <td>
                    <a href="#" :data-key="customer.id" @click="deleteCustomer">Remover</a>
                </td>
            </tr>
        </tbody>
    </table>
</template>

<script>

export default {
    data() {
        return {
            customers: []
        }
    },
    methods: {
        initCustomers: function() {
            fetch("http://localhost:8081/customers/").then(resp => resp.json())
            .then(data => this.customers = data)
            .catch(e => { 
                alert(e)
                this.customers = [];
            })
        },
        deleteCustomer: function (e) {
            e.preventDefault()
            const id = e.target.getAttribute("data-key")
            fetch(`http://localhost:8081/customers/${id}`, {
                method: "DELETE"
            })
            .then(data => this.initCustomers())
        },
    },
    created () {
        this.initCustomers()
    }
}
</script>

<style>
    .listCustomers {
        margin-top: 14px;
    }
    .listCustomers thead {
        background: #ccc;
        color: white;
    }
</style>
