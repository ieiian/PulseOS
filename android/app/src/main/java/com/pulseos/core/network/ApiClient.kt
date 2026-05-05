package com.pulseos.core.network

import com.google.gson.Gson
import com.google.gson.reflect.TypeToken
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.withContext
import okhttp3.MediaType.Companion.toMediaType
import okhttp3.OkHttpClient
import okhttp3.Request
import okhttp3.RequestBody.Companion.toRequestBody
import java.io.IOException

class ApiException(message: String) : Exception(message)

class ApiClient(
    private val baseUrl: String,
    private val client: OkHttpClient = OkHttpClient(),
    private val gson: Gson = Gson(),
) {
    private val jsonMediaType = "application/json; charset=utf-8".toMediaType()

    suspend fun <T> get(path: String, type: TypeToken<T>): T = withContext(Dispatchers.IO) {
        val request = buildRequest(path).get().build()
        execute(request, type)
    }

    suspend fun <T, B> post(path: String, body: B, bodyType: TypeToken<B>, responseType: TypeToken<T>): T =
        withContext(Dispatchers.IO) {
            val json = gson.toJson(body, bodyType.type)
            val request = buildRequest(path).post(json.toRequestBody(jsonMediaType)).build()
            execute(request, responseType)
        }

    suspend fun <T, B> put(path: String, body: B, bodyType: TypeToken<B>, responseType: TypeToken<T>): T =
        withContext(Dispatchers.IO) {
            val json = gson.toJson(body, bodyType.type)
            val request = buildRequest(path).put(json.toRequestBody(jsonMediaType)).build()
            execute(request, responseType)
        }

    suspend fun <T> postEmpty(path: String, responseType: TypeToken<T>): T = withContext(Dispatchers.IO) {
        val request = buildRequest(path).post("".toRequestBody(jsonMediaType)).build()
        execute(request, responseType)
    }

    private fun buildRequest(path: String): Request.Builder {
        val url = baseUrl.trimEnd('/') + path
        return Request.Builder().url(url).header("Content-Type", "application/json")
    }

    private fun <T> execute(request: Request, type: TypeToken<T>): T {
        val response = try {
            client.newCall(request).execute()
        } catch (e: IOException) {
            throw ApiException("网络错误: ${e.message}")
        }

        if (!response.isSuccessful) {
            throw ApiException("服务器错误 (${response.code})")
        }

        val body = response.body?.string() ?: throw ApiException("空响应")
        return try {
            gson.fromJson<T>(body, type.type)
        } catch (e: Exception) {
            throw ApiException("数据解析失败: ${e.message}")
        }
    }
}
