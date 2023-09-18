import express from 'express';
import { User } from './models/User';
import userRoutes from './routes/userRoutes';

const app = express();
app.use('/api', userRoutes);
const PORT = process.env.PORT || 3000;

const mockDatabase: User[] = [];

// Set up middleware, routes, and authentication here

app.listen(PORT, () => {
    console.log(`Server is running on port ${PORT}`);
});
