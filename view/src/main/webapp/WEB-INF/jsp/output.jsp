<%@ taglib prefix="c" uri="http://java.sun.com/jsp/jstl/core" %>
<html>
<body>
    <h2>Here are the amazing words that have been entered!</h2>
    <table border="1">
        <c:forEach var="word" items="${wordList}">
            <tr>
                <td>${word.word}</td><td>${word.count}</td>
            </tr>
        </c:forEach>
    </table>
</body>
</html>