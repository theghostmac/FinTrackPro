import { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';

interface User {
    UserID: string;
    UserName: string;
    Email: string;
    Transactions: object,
}

const Dashboard = () => {
    const [userData, setUserData] = useState<User | null>(null);

    useEffect(() => {
        const fetchProfile = () => {
            const token = localStorage.getItem('userToken');
            fetch('http://localhost:9020/profile', {
                headers: {
                    Authorization: `Bearer ${token}`,
                },
            })
                .then(response => {
                    if (!response.ok) {
                        throw new Error('Error fetching profile');
                    }
                    return response.json();
                })
                .then(data => {
                    setUserData(data);
                })
                .catch(error => {
                    console.error('There was an error fetching the profile:', error);
                });
        };

        fetchProfile();
    }, []);

    return (
        <div>
            <h1>Dashboard</h1>
            <p>Welcome back, {userData?.UserName}! Here's your financial overview.</p>
            {/* Display transactions and other user data */}
            <div className="actions">
                <Link to="/create-transaction">New Transaction</Link>
                <Link to="/create-budget">New Budget</Link>
            </div>
        </div>
    );
}

export default  Dashboard;