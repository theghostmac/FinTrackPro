import React, { useState, ChangeEvent, FormEvent } from 'react';

const Register = () => {
    const [form, setForm] = useState({
        userName: '',
        email: '',
        password: '',
    });

    const handleChange = (e: ChangeEvent<HTMLInputElement>) => {
        const { name, value } = e.target;
        setForm(prevForm => ({ ...prevForm, [name]: value }));
    };

    const handleSubmit = async (e: FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        try {
            const response = await fetch('http://localhost:9020/register', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(form),
            });

            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            console.log('User registered successfully');
            // Handle successful registration, e.g., redirect to login page
        } catch (error) {
            console.error('There was an error registering the user:', error);
        }
    };

    return (
        <div className="register-container">
            <h2>Register</h2>
            <form onSubmit={handleSubmit}>
                <label htmlFor="userName">Username</label>
                <input
                    id="userName"
                    type="text"
                    name="userName"
                    placeholder="Username"
                    value={form.userName}
                    onChange={handleChange}
                />

                <label htmlFor="email">Email</label>
                <input
                    id="email"
                    type="email"
                    name="email"
                    placeholder="Email"
                    value={form.email}
                    onChange={handleChange}
                />

                <label htmlFor="password">Password</label>
                <input
                    id="password"
                    type="password"
                    name="password"
                    placeholder="Password"
                    value={form.password}
                    onChange={handleChange}
                />

                <button type="submit">Register</button>
            </form>
        </div>
    );
};

export default Register;
