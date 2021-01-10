export interface AuthValidator {
    validateToken(token: string): ParsedToken;
}

export class JwtAuthValidator implements AuthValidator {
    validateToken(token: string): ParsedToken {
        throw new Error('Method not implemented.');
    }
}

export type Role = 'initiator' | 'responder';

export interface ParsedToken {
    responderId: string;
    role: Role;
    expiry: number;
}
