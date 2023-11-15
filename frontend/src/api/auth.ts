export const Login = async (username: string, password: string) => {
    const result = await fetch("http://127.0.0.1:3000/users/login", {
        method: "POST",
        body: JSON.stringify({
            username: username,
            password: password,
        })
    }).then(res => res.json())

    console.log(result)
}