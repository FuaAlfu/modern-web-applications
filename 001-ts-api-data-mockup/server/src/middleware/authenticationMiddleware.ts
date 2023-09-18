import { Request, Response, NextFunction } from 'express';

export const authenticationMiddleware = (req: Request, res: Response, next: NextFunction) => {
    // Implement authentication logic here
    const isAuthenticated = true; //need changed
    if (!isAuthenticated) {
        return res.status(401).json({ message: 'Unauthorized' });
    }
    next();
};
