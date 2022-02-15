import React, {useState} from 'react';
import './LoginPage.scss'

export default function LoginPage({sentUsername}) {
  const [username, setUsername] = useState('')

  const handleSubmit = (event) =>{
    event.preventDefault();
    sentUsername(username)
    console.log(username)
  }
  return (
    <div className='LoginPage'>
    <form onSubmit={handleSubmit}>
    <h2>Enter Username</h2>
      <label>
        <input type="text" name="name" value = {username} onChange = {(e) => setUsername(e.target.value)} />
      </label>
      <input type="submit" value="Submit"/>
    </form>
  </div>
  );
}
