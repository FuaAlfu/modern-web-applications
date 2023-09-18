import express from 'express';
import UserController from '../controllers/UserController';

const router = express.Router();

// Set up routes for the UserController
router.use('/users', UserController);

export default router;
