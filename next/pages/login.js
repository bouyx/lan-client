
export default function Login() {
  return (
    <form action="/">

      <div >
        <label><b>Username</b></label>
        <input type="text" placeholder="Enter Username" name="uname" />
        <label><b>Password</b></label>
        <input type="password" placeholder="Enter Password" name="psw" />
        <button type="submit">Login</button>
      </div>
    </form>
  )
}