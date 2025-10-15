import { useState } from 'react'
import { useForm } from 'react-hook-form'
import './App.css'
import axios from 'axios'

function App() {

  const { register, handleSubmit } = useForm()
  const onSubmit = (data) => {
    console.log("Submit Btn Clicked !!")
    console.log("Data: ", data)

    createUser(data)
  }

  const createUser = async (data) => {
    const Data = new FormData;
    Data.append("firstName", data.firstname)
    Data.append("lastname", data.lastname)
    Data.append("email", data.email)
    Data.append("username", data.username)
    Data.append("password", data.password)

    await axios.post("http://localhost:8080/create", Data, {
      headers: {
        // 'Content-Type': 'multipart/form-data'
        'Content-Type': 'application/json'
      }
    })
  }

  return (
    <div className="outerDiv flex justify-center items-center h-screen w-screen">
      <form
        onSubmit={handleSubmit(onSubmit)}
        className="registrationBox bg-white p-8 rounded-lg shadow-md w-full max-w-md"
      >
        <h2 className="text-2xl font-bold mb-6 text-center">User Registration</h2>
        <div className="divider"></div>

        <div className="mb-4">
          <label htmlFor="firstname" className="block mb-1 font-medium">First Name</label>
          <input
            id="firstname"
            type="text"
            {...register('firstname', { required: true })}
            className="w-full px-3 py-2 border rounded focus:outline-blue-500"
            placeholder="Enter your first name"
          />
        </div>

        <div className="mb-4">
          <label htmlFor="lastname" className="block mb-1 font-medium">Last Name</label>
          <input
            id="lastname"
            type="text"
            {...register('lastname', { required: true })}
            className=" w-full px-3 py-2 border rounded focus:outline-blue-500"
            placeholder="Enter your last name"
          />
        </div>

        <div className="mb-4">
          <label htmlFor="email" className="block mb-1 font-medium">Email</label>
          <input
            id="email"
            type="email"
            {...register('email', { required: true })}
            className=" w-full px-3 py-2 border rounded focus:outline-blue-500"
            placeholder="Enter your email"
          />
        </div>

        <div className="mb-4">
          <label htmlFor="username" className="block mb-1 font-medium">Username</label>
          <input
            id="username"
            type="text"
            {...register('username', { required: true })}
            className=" w-full px-3 py-2 border rounded focus:outline-blue-500"
            placeholder="Choose a username"
          />
        </div>

        <div className="mb-6">
          <label htmlFor="password" className="block mb-1 font-medium">Password</label>
          <input
            id="password"
            type="password"
            {...register('password', { required: true })}
            className=" w-full px-3 py-2 border rounded focus:outline-blue-500"
            placeholder="Create a password"
          />
        </div>

        <div className="submitBtn flex justify-center">
          <button
            type="submit"
            className="bg-blue-600 text-white px-6 py-2 rounded hover:bg-blue-700 transition"
          >
            Submit
          </button>
        </div>
      </form>
    </div>
  )
}

export default App
