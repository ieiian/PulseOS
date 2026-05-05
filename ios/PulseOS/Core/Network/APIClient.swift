import Foundation

enum APIError: LocalizedError {
    case invalidURL
    case invalidResponse(statusCode: Int)
    case decoding(Error)
    case network(Error)

    var errorDescription: String? {
        switch self {
        case .invalidURL:
            return "无效的请求地址"
        case .invalidResponse(let code):
            return "服务器错误 (\(code))"
        case .decoding(let error):
            return "数据解析失败: \(error.localizedDescription)"
        case .network(let error):
            return "网络错误: \(error.localizedDescription)"
        }
    }
}

actor APIClient {
    private let baseURL: URL
    private let session: URLSession

    init(baseURL: URL = AppConfig.apiBaseURL, session: URLSession = .shared) {
        self.baseURL = baseURL
        self.session = session
    }

    func get<T: Decodable>(_ path: String, as type: T.Type) async throws -> T {
        let request = try buildRequest(path: path, method: "GET")
        return try await execute(request, as: type)
    }

    func post<T: Decodable, B: Encodable>(_ path: String, body: B) async throws -> T {
        var request = try buildRequest(path: path, method: "POST")
        request.httpBody = try JSONEncoder().encode(body)
        return try await execute(request, as: T.self)
    }

    func put<T: Decodable, B: Encodable>(_ path: String, body: B) async throws -> T {
        var request = try buildRequest(path: path, method: "PUT")
        request.httpBody = try JSONEncoder().encode(body)
        return try await execute(request, as: T.self)
    }

    func postEmpty<T: Decodable>(_ path: String) async throws -> T {
        var request = try buildRequest(path: path, method: "POST")
        request.httpBody = Data()
        return try await execute(request, as: T.self)
    }

    private func buildRequest(path: String, method: String) throws -> URLRequest {
        guard let url = URL(string: path, relativeTo: baseURL) else {
            throw APIError.invalidURL
        }
        var request = URLRequest(url: url)
        request.httpMethod = method
        request.setValue("application/json", forHTTPHeaderField: "Content-Type")
        return request
    }

    private func execute<T: Decodable>(_ request: URLRequest, as type: T.Type) async throws -> T {
        do {
            let (data, response) = try await session.data(for: request)
            guard let http = response as? HTTPURLResponse else {
                throw APIError.invalidResponse(statusCode: -1)
            }
            guard (200...299).contains(http.statusCode) else {
                throw APIError.invalidResponse(statusCode: http.statusCode)
            }
            do {
                return try JSONDecoder().decode(T.self, from: data)
            } catch {
                throw APIError.decoding(error)
            }
        } catch let error as APIError {
            throw error
        } catch {
            throw APIError.network(error)
        }
    }
}
