import { useState, ChangeEvent, FormEvent } from "react";
import { useNavigate} from 'react-router-dom';

const Login = () => {
    const [form, setForm] = useState({ email: '', password: ''});
    const navigate = useNavigate();

    const handleChange = (e: ChangeEvent<HTMLInputElement>) => {
        const { name, value } = e.target;
        setForm(prevForm => ({ ...prevForm, [name]: value }));
    };

    const handleSubmit = async (e: FormEvent<HTMLInputElement>) => {
        e.preventDefault();
        try {
            const response = await fetch('http://localhost:9020/login', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(form),
            });

            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            const data = await response.json();
            console.log('User logged in successfully');
            localStorage.setItem('token', data.token);
            navigate('/dashboard')
        } catch (error) {
            console.error('There was an error logging in: ', error);
        }
    };

    return (
        <div className="login-container">
            <h2>Login</h2>
            <form onSubmit={handleSubmit}>
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

                <button type="submit">Login</button>
            </form>
        </div>
    );
}

export default Login;