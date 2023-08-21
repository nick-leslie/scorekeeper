<script>
    let i = 1;
    function increment() {
        i+=1;
    }
</script>
<body>
    <h1>{i}</h1>
    <button on:click={increment}> click to increase </button>
</body>