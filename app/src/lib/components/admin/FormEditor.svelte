<script lang="ts">
	import { goto } from '$app/navigation';
	import { onMount } from 'svelte';
	import { getResource, createResource, updateResource } from '$lib/api/***REMOVED***';
	import ImageUpload from '$lib/components/***REMOVED***/ImageUpload.svelte';

	interface FieldDef {
		key: string;
		label: string;
		type?: 'text' | 'textarea' | 'json' | 'boolean' | 'select' | 'number' | 'image';
		required?: boolean;
		options?: string[];
		placeholder?: string;
	}

	let {
		resource,
		id = null,
		fields,
		title,
		backHref
	}: {
		resource: string;
		id?: string | null;
		fields: FieldDef[];
		title: string;
		backHref: string;
	} = $props();

	let formData = $state<Record<string, unknown>>({});
	let loading = $state(false);
	let saving = $state(false);
	let error = $state('');

	const isEdit = $derived(!!id);

	onMount(async () => {
		if (id) {
			loading = true;
			try {
				formData = await getResource(resource, id);
			} catch (e) {
				error = e instanceof Error ? e.message : 'Failed to load record';
			}
			loading = false;
		} else {
			// Set defaults
			for (const field of fields) {
				if (field.type === 'boolean') formData[field.key] = false;
				else if (field.type === 'json') formData[field.key] = [];
				else if (field.type === 'number') formData[field.key] = 0;
				else formData[field.key] = '';
			}
			// Default is_published to true for new records
			if ('is_published' in formData || fields.some(f => f.key === 'is_published')) {
				formData['is_published'] = true;
			}
		}
	});

	async function handleSubmit(e: SubmitEvent) {
		e.preventDefault();
		saving = true;
		error = '';

		try {
			// Clean up data before sending
			const payload: Record<string, unknown> = {};
			for (const field of fields) {
				let val = formData[field.key];
				if (field.type === 'json' && typeof val === 'string') {
					try { val = JSON.parse(val); } catch { val = []; }
				}
				if (field.type === 'number' && typeof val === 'string') {
					val = Number(val);
				}
				payload[field.key] = val;
			}

			if (isEdit && id) {
				await updateResource(resource, id, payload);
			} else {
				await createResource(resource, payload);
			}
			goto(backHref);
		} catch (e) {
			error = e instanceof Error ? e.message : 'Save failed';
		}
		saving = false;
	}

	function getJsonString(val: unknown): string {
		if (typeof val === 'string') return val;
		if (val === null || val === undefined) return '[]';
		return JSON.stringify(val, null, 2);
	}
</script>

<div class="***REMOVED***-header">
	<h1>{isEdit ? 'Edit' : 'New'} {title}</h1>
	<a href={backHref} class="***REMOVED***-btn">← Back</a>
</div>

{#if loading}
	<p style="color: var(--***REMOVED***-text-muted);">Loading...</p>
{:else}
	{#if error}
		<div class="***REMOVED***-error">{error}</div>
	{/if}

	<div class="***REMOVED***-card">
		<form class="***REMOVED***-form" onsubmit={handleSubmit}>
			{#each fields as field}
				<div class="***REMOVED***-field">
					<label for={field.key}>{field.label}{field.required ? ' *' : ''}</label>

					{#if field.type === 'textarea'}
						<textarea
							id={field.key}
							class="***REMOVED***-input ***REMOVED***-textarea"
							bind:value={formData[field.key] as string}
							required={field.required}
							placeholder={field.placeholder}
						></textarea>
					{:else if field.type === 'json'}
						<textarea
							id={field.key}
							class="***REMOVED***-input ***REMOVED***-textarea"
							value={getJsonString(formData[field.key])}
							oninput={(e) => { formData[field.key] = (e.target as HTMLTextAreaElement).value; }}
							placeholder='["item1", "item2"]'
							style="font-family: var(--***REMOVED***-mono); font-size: 0.8rem;"
						></textarea>
					{:else if field.type === 'boolean'}
						<label style="display:flex; align-items:center; gap:0.5rem; cursor:pointer;">
							<input type="checkbox" bind:checked={formData[field.key] as boolean} />
							<span style="font-size:0.85rem;">{formData[field.key] ? 'Yes' : 'No'}</span>
						</label>
					{:else if field.type === 'select' && field.options}
						<select id={field.key} class="***REMOVED***-input ***REMOVED***-select" bind:value={formData[field.key] as string}>
							{#each field.options as opt}
								<option value={opt}>{opt}</option>
							{/each}
						</select>
					{:else if field.type === 'number'}
						<input
							id={field.key}
							class="***REMOVED***-input"
							type="number"
							bind:value={formData[field.key] as number}
							required={field.required}
						/>
					{:else if field.type === 'image'}
						<ImageUpload
							value={String(formData[field.key] || '')}
							onchange={(url) => { formData[field.key] = url; }}
						/>
					{:else}
						<input
							id={field.key}
							class="***REMOVED***-input"
							type="text"
							bind:value={formData[field.key] as string}
							required={field.required}
							placeholder={field.placeholder}
						/>
					{/if}
				</div>
			{/each}

			<div style="display:flex; gap:0.75rem; margin-top:0.5rem;">
				<button class="***REMOVED***-btn ***REMOVED***-btn-primary" type="submit" disabled={saving}>
					{saving ? 'Saving...' : isEdit ? 'Update' : 'Create'}
				</button>
				<a href={backHref} class="***REMOVED***-btn">Cancel</a>
			</div>
		</form>
	</div>
{/if}
