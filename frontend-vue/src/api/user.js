import { API_BASE_URL } from './config';

export function login(username, password) {
  return fetch(`${API_BASE_URL}/login`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ username, password })
  }).then(res => res.json());
}

export function register(username, password) {
  return fetch(`${API_BASE_URL}/register`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ username, password })
  }).then(res => res.json());
}
