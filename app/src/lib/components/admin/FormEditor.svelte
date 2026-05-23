<script lang="ts">
	import { goto } from '$app/navigation';
	import { onMount } from 'svelte';
	import { getResource, createResource, updateResource } from '$lib/api/***REMOVED***';
	import ImageUpload from '$lib/components/***REMOVED***/ImageUpload.svelte';
	import TagListEditor from '$lib/components/***REMOVED***/TagListEditor.svelte';
	import GalleryEditor from '$lib/components/***REMOVED***/GalleryEditor.svelte';

	interface FieldDef {
		key: string;
		label: string;
		type?: 'text' | 'textarea' | 'json' | 'tags' | 'gallery' | 'boolean' | 'select' | 'number' | 'image';
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
			for (const field of fields) {
				if (field.type === 'boolean') formData[field.key] = false;
				else if (field.type === 'json' || field.type === 'tags' || field.type === 'gallery') formData[field.key] = [];
				else if (field.type === 'number') formData[field.key] = 0;
				else formData[field.key] = '';
			}
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

	function getArrayValue(val: unknown): string[] {
		if (Array.isArray(val)) return val as string[];
		if (typeof val === 'string') {
			try { return JSON.parse(val); } catch { return []; }
		}
		return [];
	}

	function getGalleryValue(val: unknown): { image: string; caption: string; description: string }[] {
		if (Array.isArray(val)) return val as { image: string; caption: string; description: string }[];
		if (typeof val === 'string') {
			try { return JSON.parse(val); } catch { return []; }
		}
		return [];
	}
</script>

<div class="***REMOVED***-header">
	<h1>{isEdit ? 'Edit' : 'New'} {title}</h1>
	<a href={backHref} class="***REMOVED***-btn">← Back</a>
</div>

{#if loading}
	<div class="form-loading">
		<div class="loading-box"></div>
		<span>Loading record...</span>
	</div>
{:else}
	{#if error}
		<div class="***REMOVED***-error">{error}</div>
	{/if}

	<div class="***REMOVED***-card form-card">
		<div class="form-stamp" aria-hidden="true">{isEdit ? 'EDIT' : 'NEW'}</div>
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
					{:else if field.type === 'tags'}
						<TagListEditor
							value={getArrayValue(formData[field.key])}
							onchange={(items) => { formData[field.key] = items; }}
							placeholder={field.placeholder || 'Add item...'}
						/>
					{:else if field.type === 'gallery'}
						<GalleryEditor
							value={getGalleryValue(formData[field.key])}
							onchange={(items) => { formData[field.key] = items; }}
						/>
					{:else if field.type === 'json'}
						<textarea
							id={field.key}
							class="***REMOVED***-input ***REMOVED***-textarea json-field"
							value={getJsonString(formData[field.key])}
							oninput={(e) => { formData[field.key] = (e.target as HTMLTextAreaElement).value; }}
							placeholder='["item1", "item2"]'
						></textarea>
					{:else if field.type === 'boolean'}
						<label class="bool-toggle">
							<input type="checkbox" bind:checked={formData[field.key] as boolean} />
							<span class="bool-label">{formData[field.key] ? 'Yes' : 'No'}</span>
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

			<div class="form-actions">
				<button class="***REMOVED***-btn ***REMOVED***-btn-primary" type="submit" disabled={saving}>
					{saving ? '◈ Saving...' : isEdit ? '→ Update' : '→ Create'}
				</button>
				<a href={backHref} class="***REMOVED***-btn">Cancel</a>
			</div>
		</form>
	</div>
{/if}

<style>
	.form-loading {
		display: flex;
		align-items: center;
		gap: 0.75rem;
		font-family: var(--font-mono);
		font-size: 0.75rem;
		text-transform: uppercase;
		color: var(--***REMOVED***-ink-muted);
	}

	.loading-box {
		width: 14px;
		height: 14px;
		border: 3px solid var(--***REMOVED***-ink);
		animation: spin-box 0.8s linear infinite;
	}

	@keyframes spin-box {
		from { transform: rotate(0deg); }
		to { transform: rotate(360deg); }
	}

	.form-card {
		position: relative;
	}

	.form-stamp {
		position: absolute;
		top: -8px;
		right: 16px;
		font-family: var(--font-mono);
		font-size: 0.58rem;
		font-weight: 700;
		text-transform: uppercase;
		letter-spacing: 1.5px;
		background: var(--***REMOVED***-blue);
		border: 3px solid var(--***REMOVED***-ink);
		box-shadow: 3px 3px 0 0 var(--***REMOVED***-ink);
		padding: 0.15rem 0.45rem;
	}

	.json-field {
		font-family: var(--font-mono);
		font-size: 0.78rem;
	}

	.bool-toggle {
		display: flex;
		align-items: center;
		gap: 0.55rem;
		cursor: pointer;
	}

	.bool-toggle input {
		width: 18px;
		height: 18px;
		accent-color: var(--***REMOVED***-ink);
		cursor: pointer;
	}

	.bool-label {
		font-family: var(--font-mono);
		font-size: 0.78rem;
		text-transform: uppercase;
	}

	.form-actions {
		display: flex;
		gap: 0.75rem;
		margin-top: 0.5rem;
		padding-top: 1rem;
		border-top: 2px solid var(--***REMOVED***-ink);
	}
</style>
