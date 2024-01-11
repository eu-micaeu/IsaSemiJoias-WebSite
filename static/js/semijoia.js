document.addEventListener("DOMContentLoaded", () => {
    const comprar = document.getElementById("comprar");

    comprar.addEventListener("click", () => {
        // Definir a variável com o valor 100
        let preco = 100;

        // Armazenar a variável no localStorage
        localStorage.setItem("preco", preco);

        // Redirecionar para "/pagamento"
        window.location.href = "/pagamento";
    });
});
