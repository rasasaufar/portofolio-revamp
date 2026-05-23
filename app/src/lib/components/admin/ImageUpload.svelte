<script lang="ts">
	import { get } from 'svelte/store';
	import { token } from '$lib/stores/auth';
	import { API_BASE_URL } from '$lib/config';

	let {
		value = '',
		onchange
	}: {
		value?: string;
		onchange?: (url: string) => void;
	} = $props();

	let uploading = $state(false);
	let error = $state('');
	let dragOver = $state(false);

	async function uploadFile(file: File) {
		uploading = true;
		error = '';

		const formData = new FormData();
		formData.append('file', file);

		try {
			const currentToken = get(token);
			const res = await fetch(`${API_BASE_URL}/api/upload`, {
				method: 'POST',
				headers: {
					'Authorization': `Bearer ${currentToken}`
				},
				body: formData
			});

			if (!res.ok) {
				const data = await res.json().catch(() => ({}));
				throw new Error(data.error || 'Upload failed');
			}

			const data = await res.json();
			value = data.url;
			onchange?.(data.url);
		} catch (e) {
			error = e instanceof Error ? e.message : 'Upload failed';
		} finally {
			uploading = false;
		}
	}

	function handleFileInput(e: Event) {
		const input = e.target as HTMLInputElement;
		const file = input.files?.[0];
		if (file) uploadFile(file);
	}

	function handleDrop(e: DragEvent) {
		e.preventDefault();
		dragOver = false;
		const file = e.dataTransfer?.files?.[0];
		if (file) uploadFile(file);
	}

	function handleDragOver(e: DragEvent) {
		e.preventDefault();
		dragOver = true;
	}

	function handleDragLeave() {
		dragOver = false;
	}

	function clearImage() {
		value = '';
		onchange?.('');
	}
</script>

<div class="image-upload">
	{#if value}
		<div class="image-preview">
			<img src={value} alt="Preview" />
			<div class="image-preview-actions">
				<button type="button" class="***REMOVED***-btn ***REMOVED***-btn-sm ***REMOVED***-btn-danger" onclick={clearImage}>Remove</button>
			</div>
		</div>
		<input type="text" class="***REMOVED***-input" {value} readonly style="margin-top:0.5rem; font-size:0.75rem; color:var(--***REMOVED***-text-muted);" />
	{:else}
		<div
			class="upload-dropzone"
			class:drag-over={dragOver}
			ondrop={handleDrop}
			ondragover={handleDragOver}
			ondragleave={handleDragLeave}
			role="button"
			tabindex="0"
		>
			{#if uploading}
				<p>Uploading...</p>
			{:else}
				<p>📁 Drag & drop image here</p>
				<p style="font-size:0.75rem; color:var(--***REMOVED***-text-muted);">or click to browse</p>
				<input type="file" accept="image/*" onchange={handleFileInput} class="file-input" />
			{/if}
		</div>
	{/if}

	{#if error}
		<p class="upload-error">{error}</p>
	{/if}
</div>

<style>
	.image-upload {
		width: 100%;
	}

	.image-preview {
		position: relative;
		border: 1px solid var(--***REMOVED***-border);
		border-radius: var(--***REMOVED***-radius);
		overflow: hidden;
		max-width: 300px;
	}

	.image-preview img {
		width: 100%;
		height: 160px;
		object-fit: cover;
		display: block;
	}

	.image-preview-actions {
		position: absolute;
		top: 0.5rem;
		right: 0.5rem;
	}

	.upload-dropzone {
		position: relative;
		border: 2px dashed var(--***REMOVED***-border);
		border-radius: var(--***REMOVED***-radius);
		padding: 2rem;
		text-align: center;
		cursor: pointer;
		transition: border-color 0.15s, background 0.15s;
	}

	.upload-dropzone:hover,
	.upload-dropzone.drag-over {
		border-color: var(--***REMOVED***-accent);
		background: rgba(99, 102, 241, 0.05);
	}

	.upload-dropzone p {
		margin: 0.25rem 0;
		color: var(--***REMOVED***-text-muted);
	}

	.file-input {
		position: absolute;
		inset: 0;
		opacity: 0;
		cursor: pointer;
	}

	.upload-error {
		color: var(--***REMOVED***-danger);
		font-size: 0.8rem;
		margin-top: 0.4rem;
	}
</style>
