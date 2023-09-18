import { Request, Response } from 'express';
import { User } from '../models/User';
import { authenticationMiddleware } from '../middleware/authenticationMiddleware';

const router = express.Router();

// Protect all routes in this controller with authentication middleware
router.use(authenticationMiddleware);

// Implement CRUD actions here
router.get('/', (req: Request, res: Response) => {
    // Retrieve all users from the mock database
    res.json(mockDatabase);
});

router.get('/:id', (req: Request, res: Response) => {
    // Retrieve a user by ID from the mock database
    const userId = parseInt(req.params.id);
    const user = mockDatabase.find(u => u.id === userId);
    if (!user) {
        return res.status(404).json({ message: 'User not found' });
    }
    res.json(user);
});

// Implement other CRUD routes (POST, PUT, DELETE)

export default router;
