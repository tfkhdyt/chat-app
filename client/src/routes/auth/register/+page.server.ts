import { message, superValidate } from 'sveltekit-superforms';
import { zod } from 'sveltekit-superforms/adapters';
import { registerSchema } from '$lib/schema/auth';
import { fail, type Actions } from '@sveltejs/kit';
import { API_URL } from '$env/static/private';

export const load = async () => {
	const form = await superValidate(zod(registerSchema));
	return { form };
};

export const actions = {
	default: async ({ request }) => {
		const form = await superValidate(request, zod(registerSchema));
		if (!form.valid) {
			return fail(400, { form });
		}

		try {
			const response = await fetch(`${API_URL}/auth/register`, {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify({
					username: form.data.username,
					email: form.data.email,
					password: form.data.password
				})
			});
			const data = await response.json();
			if (!response.ok || !data.success) {
				return message(form, { type: 'error', text: data.message });
			}

			return message(form, { type: 'success', text: 'Account created successfully' });
		} catch (error) {
			console.error(error);

			if (error instanceof Error) {
				return message(form, { type: 'error', text: error.message }, { status: 400 });
			}

			return message(
				form,
				{ type: 'error', text: `An unknown error occurred, ${String(error)}` },
				{ status: 500 }
			);
		}
	}
} satisfies Actions;
