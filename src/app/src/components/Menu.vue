<template>
    <nav class="navbar">
        <h1>Sistema de cadastro</h1>
        <ul>
            <li><a href="#/">Home</a></li>
            <li><a href="#/novo">Novo</a></li>
            <li><a href="#/lista">Lista</a></li>
        </ul>
    </nav>
</template>

<script>
export default {

}
</script>

<style>
    .navbar {
        display: flex;
        justify-content: space-around;
        margin-top: 15px;
        width: 100%;
    }
    .navbar ul {
        display: flex;
        width: 30%;
        list-style-type: none;
    }
    .navbar ul li{
        flex: 1;
    }
    .navbar ul li a {
        text-decoration: none;
    }
</style>
